(() => {
  window.open(`URL_ROOT/api/stash?url=${encodeURIComponent(window.location)}&key=${encodeURIComponent('STASH_KEY')}`)
})();
