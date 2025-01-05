import { createContext } from "react";
import { Profile } from "../../../services/auth/types";

export const User = createContext<Profile | undefined>(undefined);
