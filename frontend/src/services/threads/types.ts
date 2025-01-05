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
  Deleted: boolean;
}

export interface ThreadTag {
  ThreadID: number;
  Tag: string;
}
