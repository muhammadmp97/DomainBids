const timeForHuman = function (time) {
  let diff = Math.floor((new Date() - time) / 1000)

  if (diff < 5) {
    return `now`
  }

  if (diff < 60) {
    return `${diff}s ago`
  }

  if (diff < 3600) {
    diff = Math.floor(diff / 60)
    return `${diff}m ago`
  }

  if (diff < 3600 * 24) {
    diff = Math.floor(diff / 3600)
    return `${diff}h ago`
  }

  diff = Math.floor(diff / (3600 * 24))
  return `${diff}d ago`
}

const isFuture = function (date) {
  const now = new Date()

  return date - now > 0
}

export { timeForHuman, isFuture }
