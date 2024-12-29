export interface Thread {
  ID: number;
  Title: string;
  Description: string | null;
  Body: string;
  UserID: number;
  Likes: number;
  Comments: number;
  Shares: number;
  CreatedAt: string;
  UpdatedAt: string;
  Liked: boolean | null;
  ThreadTags: ThreadTag[];
  Username: string;
  Image: string | null;
  ProfileImage: string | null;
}

export interface ThreadTag {
  ThreadID: number;
  Tag: string;
}

export interface ThreadComment {
  ID: number;
  Comment: string;
  ThreadID: string | null;
  UserID: number;
  Likes: number;
  Comments: number;
  CreatedAt: string;
  UpdatedAt: string;
  Liked: boolean | null;
  Username: string;
  ProfileImage: string;
}

export interface Profile {
  UserID: number;
  Username: string;
  Posts: number;
  Comments: number;
  Aura: number;
  CreatedAt: string;
  ProfileImage: string;
}
