import { websocketBaseURL } from "..";

/**
 *
 * @returns a Websocket connection
 */
export function getNotifications(): WebSocket {
  return new WebSocket(`${websocketBaseURL}/notifications`);
}

/**
 *
 * @param id the id of the thread
 * @returns a Websocket connection
 */
export function getThreadInfo(id: number): WebSocket {
  return new WebSocket(`${websocketBaseURL}/thread-info/${id}`);
}
