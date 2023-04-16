/** @type {import('./$types').PageLoad} */
export const load = async ({ fetch }) => {
  const response = await fetch('/api/hello')
  const data = await response.json()

  return data
}
