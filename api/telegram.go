package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/telegram"
	"net/http"
)

func (a *API) telegram(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Telegram-Bot-Api-Secret-Token") != a.runtime.Env.TelegramSecretToken {
		errorResponse(w, http.StatusForbidden, errors.New("no token"))
		return
	}

	bodybytes, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	var update telegram.Update
	err = json.Unmarshal(bodybytes, &update)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	if fmt.Sprint(update.Message.From.Id) != a.runtime.Env.TelegramAllowedId {
		errorResponse(w, http.StatusForbidden, errors.New("user not permitted"))
		return
	}

	item, err := a.process(update.Message.Message.Text)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	response := telegram.OutgoingMessage{
		ChatId: update.Message.Chat.Id,
		Message: telegram.Message{
			Text: item.Name,
		},
	}

	err = telegram.CallMethod(a.runtime.Env.TelegramAPIToken, "sendMessage", response)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
