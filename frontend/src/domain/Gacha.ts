export enum Rarity {
  Common = "Common",
  Rare = "Rare",
  Epic = "Epic",
  Legendary = "Legendary"
}

export interface GachaResult {
  id: number;
  itemName: string;
  rarity: Rarity;
  pointsEarned: number;
}

export interface GachaHistory extends GachaResult {
  createdAt: string;
}