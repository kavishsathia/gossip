import { Profile, Thread, ThreadComment } from "./types";

// Update both
const baseURL =
  "https://gossip.s6wyfaw6z9q0r.ap-southeast-1.cs.amazonlightsail.com";
const websocketBaseURL =
  "wss://gossip.s6wyfaw6z9q0r.ap-southeast-1.cs.amazonlightsail.com";

const request = async (input: RequestInfo | URL, init?: RequestInit) => {
  const response = await fetch(input, { ...init, credentials: "include" });
  if (response.status === 401) {
    window.location.href = "/login";
  }
  return response;
};

export async function listThreads(query: string): Promise<Thread[]> {
  const response = await request(`${baseURL}/threads?query=${query}`);
  const json = await response.json();
  return json;
}

export async function getThread(id: number): Promise<Thread> {
  const response = await request(`${baseURL}/thread/${id}`);
  const json = await response.json();
  return json;
}

export async function likeThread(id: number): Promise<Thread> {
  const response = await request(`${baseURL}/thread/${id}/like`, {
    method: "POST",
  });
  const json = await response.json();
  return json;
}

export async function unlikeThread(id: number): Promise<Thread> {
  const response = await request(`${baseURL}/thread/${id}/like`, {
    method: "DELETE",
  });
  const json = await response.json();
  return json;
}

export async function postThread(
  title: string,
  body: string,
  tags: string[],
  image: string
): Promise<number> {
  const response = await request(`${baseURL}/thread`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title,
      body,
      tags,
      image,
    }),
  });
  const t = await response.json();
  return t.ThreadID;
}

export async function listThreadComments(id: number): Promise<ThreadComment[]> {
  const response = await request(`${baseURL}/thread/${id}/comments`);
  const json = await response.json();
  return json;
}

export async function listThreadCommentComments(
  id: number
): Promise<ThreadComment[]> {
  const response = await request(`${baseURL}/comment/${id}/comments`);
  const json = await response.json();
  return json;
}

export async function createThreadComment(
  id: number,
  body: string
): Promise<boolean> {
  const response = await request(`${baseURL}/thread/${id}/comment`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      body,
    }),
  });
  return response.status === 200;
}

export async function createThreadCommentComment(
  id: number,
  body: string
): Promise<boolean> {
  const response = await request(`${baseURL}/comment/${id}/comment`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      body,
    }),
  });
  return response.status === 200;
}

export async function likeThreadComment(id: number): Promise<Thread> {
  const response = await request(`${baseURL}/comment/${id}/like`, {
    method: "POST",
  });
  const json = await response.json();
  return json;
}

export async function unlikeThreadComment(id: number): Promise<Thread> {
  const response = await request(`${baseURL}/comment/${id}/like`, {
    method: "DELETE",
  });
  const json = await response.json();
  return json;
}

export async function createUser(
  username: string,
  password: string
): Promise<number> {
  const response = await request(`${baseURL}/user`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      password,
    }),
  });

  if (response.status === 409) {
    return -1;
  }

  const t = await response.json();
  return t.UserID;
}

export async function loginAsUser(
  username: string,
  password: string
): Promise<boolean> {
  const response = await request(`${baseURL}/user`, {
    credentials: "include",
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      password,
    }),
  });

  return response.status === 200;
}

export async function getMe(): Promise<Profile | null> {
  const response = await request(`${baseURL}/user`, {
    method: "GET",
  });

  if (response.status != 200) {
    return null;
  }

  return await response.json();
}

export async function getUser(id: number): Promise<Profile> {
  const response = await request(`${baseURL}/user/${id}`, {
    method: "GET",
  });

  return await response.json();
}

export function getNotifications(): WebSocket {
  return new WebSocket(`${websocketBaseURL}/notifications`);
}

export function getThreadInfo(id: number): WebSocket {
  return new WebSocket(`${websocketBaseURL}/thread-info/${id}`);
}
