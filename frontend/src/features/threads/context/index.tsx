import { createContext } from "react";
import { Profile } from "../../../services/threads/types";

export const User = createContext<Profile | undefined>(undefined);
