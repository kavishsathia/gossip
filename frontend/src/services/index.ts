export const baseURL = import.meta.env.VITE_BASE_URL;
export const websocketBaseURL = import.meta.env.VITE_WEBSOCKET_BASE_URL;

export const request = async (input: RequestInfo | URL, init?: RequestInit) => {
  const response = await fetch(input, { ...init, credentials: "include" });
  if (response.status === 401) {
    if (
      confirm(`
      You are lurking, too make changes please log in. 
      Do you want to be redirected to the login page?
    `)
    )
      window.location.href = "/login";
  }

  if (response.status > 499) {
    window.location.href = `/error?status=${"Internal Server Error"}&status-text=${"My bad bruh, lemme fix that"}`;
  }
  return response;
};
