import axios from 'axios';
import type { AddProductResponse, NewProduct, Product, Status } from '../types';

const API_BASE_URL = 'https://goods-viet.onrender.com/api';

export const addProduct = async (product: NewProduct, image: File | null): Promise<AddProductResponse> => {
  const formData = new FormData();
  formData.append('name', product.name);
  formData.append('manufacturer', product.manufacturer);
  formData.append('initialStatus', product.initialStatus);
  formData.append('initialDetails', product.initialDetails);
  if (image) {
    formData.append('image', image);
  }

  const response = await axios.post(`${API_BASE_URL}/product/add`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
      'Accept': 'application/json',
    },
  });
  return response.data;
};

export const getProductById = async (id: number): Promise<Product> => {
  const response = await axios.get(`${API_BASE_URL}/product/${id}`);
  console.log(response.data);
  return response.data;
};

export const getAllProducts = async (): Promise<Product[]> => {
  const response = await axios.get(`${API_BASE_URL}/products`);
  return response.data;
};

export const updateProductStatus = async (id: number, status: Status): Promise<{ message: string }> => {
  const response = await axios.post(`${API_BASE_URL}/product/${id}/status`, status);
  return response.data;
};