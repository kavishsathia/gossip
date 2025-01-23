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
  ModerationFlag: string | null;
  ThreadCorrections: ThreadCorrection[];
  Username: string;
  Image: string | null;
  ProfileImage: string | null;
  Deleted: boolean;
}

export interface ThreadTag {
  ThreadID: number;
  Tag: string;
}

export interface ThreadCorrection {
  ThreadID: number;
  Correction: string;
}
