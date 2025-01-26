import { baseURL, request } from "..";
import { Thread } from "./types";

/**
 *
 * @param query the search string
 * @returns a list of threads
 */
export async function listThreads(query: string): Promise<Thread[]> {
  const response = await request(`${baseURL}/threads?query=${query}`);
  const json = await response.json();
  return json;
}

/**
 *
 * @param id the id of the thread
 * @returns the thread object
 */
export async function getThread(id: number): Promise<Thread> {
  const response = await request(`${baseURL}/thread/${id}`);
  const json = await response.json();
  return json;
}

/**
 *
 * @param id the id of the thread
 * @returns true if the thread was deleted
 */
export async function deleteThread(id: number) {
  const response = await request(`${baseURL}/thread/${id}`, {
    method: "DELETE",
  });
  return response.status === 200;
}

/**
 *
 * @param id the id of the thread
 * @param title the title of the thread
 * @param body the markdown body of the thread
 * @param tags the list of tags
 * @param image the image url of the thread
 * @returns a thread
 */
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

/**
 *
 * @param id the id of the thread
 * @returns true if the thread was liked
 */
export async function likeThread(id: number): Promise<boolean> {
  const response = await request(`${baseURL}/thread/${id}/like`, {
    method: "POST",
  });

  return response.status === 200;
}

/**
 *
 * @param id the id of the thread
 * @returns true if the thread was reported
 */
export async function reportThread(id: number): Promise<boolean> {
  const response = await request(`${baseURL}/thread/${id}/report`, {
    method: "PUT",
  });
  return response.status === 200;
}

/**
 *
 * @param id the id of the thread
 * @returns true if the thread was unliked
 */
export async function unlikeThread(id: number): Promise<boolean> {
  const response = await request(`${baseURL}/thread/${id}/like`, {
    method: "DELETE",
  });
  return response.status === 200;
}

/**
 *
 * @param title the title of the thread
 * @param body the markdown body of the thread
 * @param tags the list of tags
 * @param image the image url of the thread
 * @returns a thread
 */
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
