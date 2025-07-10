export interface UserPoint {
  userId: number;
  balance: number;
}

export interface PointTransaction {
  id: number;
  amount: number;
  type: "gacha" | "spend";
  description: string;
  createdAt: string;
}