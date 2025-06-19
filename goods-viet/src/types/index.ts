export interface Product {
  productId: number;
  name: string;
  manufacturer: string;
  qrCodeUrl: string;
  statuses: Status[];
  imageURL?: string;
}

export interface Status {
  status: string;
  details: string;
  timestamp: string;
}

export interface NewProduct {
  name: string;
  manufacturer: string;
  initialStatus: string;
  initialDetails: string;
}

export interface AddProductResponse {
  productId: number;
  qrCodeData: string;
  qrCodeUrl: string;
  imageUrl: string;
}