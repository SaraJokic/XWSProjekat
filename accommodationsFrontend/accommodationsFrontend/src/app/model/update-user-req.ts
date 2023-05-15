import { User } from "./user";

export interface UpdateUserReq {
    user: User,
    id: string,
    UserId: string
}
