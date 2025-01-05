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
  Deleted: boolean;
}
