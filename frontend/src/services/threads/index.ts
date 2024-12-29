import { Profile, Thread, ThreadComment } from "./types";

const request = async (input: RequestInfo | URL, init?: RequestInit) => {
  const response = await fetch(input, { ...init, credentials: "include" });
  if (response.status === 401) {
    window.location.href = "/login";
  }
  return response;
};

export async function listThreads(query: string): Promise<Thread[]> {
  const response = await request(
    `http://localhost:8080/threads?query=${query}`
  );
  const json = await response.json();
  return json;
}

export async function getThread(id: number): Promise<Thread> {
  const response = await request(`http://localhost:8080/thread/${id}`);
  const json = await response.json();
  return json;
}

export async function likeThread(id: number): Promise<Thread> {
  const response = await request(`http://localhost:8080/thread/${id}/like`, {
    method: "POST",
  });
  const json = await response.json();
  return json;
}

export async function unlikeThread(id: number): Promise<Thread> {
  const response = await request(`http://localhost:8080/thread/${id}/like`, {
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
  const response = await request("http://localhost:8080/thread", {
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
  const response = await request(`http://localhost:8080/thread/${id}/comments`);
  const json = await response.json();
  return json;
}

export async function listThreadCommentComments(
  id: number
): Promise<ThreadComment[]> {
  const response = await request(
    `http://localhost:8080/comment/${id}/comments`
  );
  const json = await response.json();
  return json;
}

export async function createThreadComment(
  id: number,
  body: string
): Promise<boolean> {
  const response = await request(`http://localhost:8080/thread/${id}/comment`, {
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
  const response = await request(
    `http://localhost:8080/comment/${id}/comment`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        body,
      }),
    }
  );
  return response.status === 200;
}

export async function likeThreadComment(id: number): Promise<Thread> {
  const response = await request(`http://localhost:8080/comment/${id}/like`, {
    method: "POST",
  });
  const json = await response.json();
  return json;
}

export async function unlikeThreadComment(id: number): Promise<Thread> {
  const response = await request(`http://localhost:8080/comment/${id}/like`, {
    method: "DELETE",
  });
  const json = await response.json();
  return json;
}

export async function createUser(
  username: string,
  password: string
): Promise<number> {
  const response = await request("http://localhost:8080/user", {
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
  const response = await request("http://localhost:8080/user", {
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
  const response = await request("http://localhost:8080/user", {
    method: "GET",
  });

  if (response.status != 200) {
    return null;
  }

  return await response.json();
}

export async function getUser(id: number): Promise<Profile> {
  const response = await request(`http://localhost:8080/user/${id}`, {
    method: "GET",
  });

  return await response.json();
}
