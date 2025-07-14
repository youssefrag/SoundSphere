export function formatDuration(totalSeconds) {
  const secs = Math.floor(totalSeconds % 60);
  const mins = Math.floor(totalSeconds / 60);
  return `${mins}:${String(secs).padStart(2, "0")}`;
}

export function formatDDMMYYYY(iso) {
  const d = new Date(iso);
  const day = String(d.getDate()).padStart(2, "0");
  const month = String(d.getMonth() + 1).padStart(2, "0");
  const year = d.getFullYear();
  return `${day}/${month}/${year}`;
}
