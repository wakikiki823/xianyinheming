export function savePlayHistory(song) {
  const history = uni.getStorageSync('playHistory') || []
  if (!history.some(item => item.id === song.id)) {
    history.unshift(song)
    uni.setStorageSync('playHistory', history.slice(0, 50))
  }
} 