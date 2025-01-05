import { baseURL, request } from "..";
import { ThreadComment } from "./types";

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

export async function likeThreadComment(id: number): Promise<Comment> {
  const response = await request(`${baseURL}/comment/${id}/like`, {
    method: "POST",
  });
  const json = await response.json();
  return json;
}

export async function unlikeThreadComment(id: number): Promise<Comment> {
  const response = await request(`${baseURL}/comment/${id}/like`, {
    method: "DELETE",
  });
  const json = await response.json();
  return json;
}

export async function deleteThreadComment(id: number): Promise<Comment> {
  const response = await request(`${baseURL}/comment/${id}`, {
    method: "DELETE",
  });
  const json = await response.json();
  return json;
}

export async function editThreadComment(
  id: number,
  body: string
): Promise<Comment> {
  const response = await request(`${baseURL}/comment/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      body,
    }),
  });
  const json = await response.json();
  return json;
}
