import { baseURL, request } from "..";
import { ThreadComment } from "./types";

/**
 *
 * @param id the id of the thread
 * @returns the list of the thread comments
 */
export async function listThreadComments(id: number): Promise<ThreadComment[]> {
  const response = await request(`${baseURL}/thread/${id}/comments`);
  const json = await response.json();
  return json;
}

/**
 *
 * @param id the id of the comment
 * @returns the list of comments
 */
export async function listThreadCommentComments(
  id: number
): Promise<ThreadComment[]> {
  const response = await request(`${baseURL}/comment/${id}/comments`);
  const json = await response.json();
  return json;
}

/**
 *
 * @param id the id of the thread
 * @param body the comment body
 * @returns true if the comment was created
 */
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

/**
 *
 * @param id the id of the comment
 * @param body the comment body
 * @returns true if the comment was created
 */
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

/**
 *
 * @param id the id of the comment
 * @returns true if the like was registered
 */
export async function likeThreadComment(id: number): Promise<boolean> {
  const response = await request(`${baseURL}/comment/${id}/like`, {
    method: "POST",
  });
  return response.status === 200;
}

/**
 *
 * @param id the id of the comment
 * @returns true if the like was removed
 */
export async function unlikeThreadComment(id: number): Promise<boolean> {
  const response = await request(`${baseURL}/comment/${id}/like`, {
    method: "DELETE",
  });
  return response.status === 200;
}

/**
 *
 * @param id the id of the comment
 * @returns true if the comment was deleted
 */
export async function deleteThreadComment(id: number): Promise<boolean> {
  const response = await request(`${baseURL}/comment/${id}`, {
    method: "DELETE",
  });
  return response.status === 200;
}

/**
 *
 * @param id the id of the comment
 * @param body the comment body
 * @returns the comment
 */
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
