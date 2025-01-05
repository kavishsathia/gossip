import { baseURL, request } from "..";
import { Thread } from "./types";

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

export async function deleteThread(id: number) {
  const response = await request(`${baseURL}/thread/${id}`, {
    method: "DELETE",
  });
  const json = await response.json();
  return json;
}

export async function editThread(
  id: number,
  title: string,
  body: string,
  tags: string[],
  image: string
) {
  const response = await request(`${baseURL}/thread/${id}`, {
    method: "PUT",
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
