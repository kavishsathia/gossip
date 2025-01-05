import { websocketBaseURL } from "..";

export function getNotifications(): WebSocket {
  return new WebSocket(`${websocketBaseURL}/notifications`);
}

export function getThreadInfo(id: number): WebSocket {
  return new WebSocket(`${websocketBaseURL}/thread-info/${id}`);
}
