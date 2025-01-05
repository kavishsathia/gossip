import { baseURL, request } from "..";
import { Profile } from "./types";

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
  const response = await request(`${baseURL}/user/sign-in`, {
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

export async function signOut(): Promise<boolean> {
  const response = await request(`${baseURL}/user/sign-out`, {
    method: "GET",
  });

  return response.status === 401;
}

export async function getMe(): Promise<Profile | null> {
  const response = await request(`${baseURL}/user/me`, {
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
