const timeForHuman = function (time) {
  let diff = Math.floor((new Date() - time) / 1000)

  if (diff < 60) {
    return `${diff}s`
  }

  if (diff < 3600) {
    diff = Math.floor(diff / 60)
    return `${diff}m`
  }

  if (diff < 3600 * 24) {
    diff = Math.floor(diff / 3600)
    return `${diff}h`
  }

  diff = Math.floor(diff / (3600 * 24))
  return `${diff}d`
}

export { timeForHuman }
