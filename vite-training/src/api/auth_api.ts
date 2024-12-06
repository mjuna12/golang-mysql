import axios from 'axios';
import { RegisterResponseTypes, LoginResponseTypes, EventResponseTypes } from '../types/auth.type';

// Register API 
export const registerApi = async (email: string, password: string): Promise<boolean> => {
  try {
    await axios.post<RegisterResponseTypes>('https://pipe.dev.dtaclass.id/api/auth/register/${id}', {
      email,
      password,
    });
    return true;
  } catch (err: any) {
    throw new Error(err.response?.data?.message || 'Registration failed');
  }
};

// Login API 
export const loginApi = async (email: string, password: string): Promise<LoginResponseTypes> => {
  try {
    const response = await axios.post<LoginResponseTypes>('https://pipe.dev.dtaclass.id/api/auth/login', {
      email,
      password,
    });
    return response.data;
  } catch (err: any) {
    throw new Error(err.response?.data?.message || 'Login failed');
  }
};

// Get Event List API 
export const getEventListApi = async (token: string): Promise<EventResponseTypes> => {
  try {
    const response = await axios.get<EventResponseTypes>('https://pipe.dev.dtaclass.id/api/events', {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data;
  } catch (err: any) {
    throw new Error('Failed to fetch events');
  }
};

// Register Event API 
export const registerEventApi = async (
  eventId: string, lat: number, lng: number, token: string
): Promise<any> => {
  try {
    const response = await axios.post(
      `https://pipe.dev.dtaclass.id/api/events/register/${eventId}`,
      { lat, lng },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (err: any) {
    throw new Error('Event registration failed');
  }
};

// Get My Events API 
export const getMyEventApi = async (token: string): Promise<any> => {
  try {
    const response = await axios.get('https://pipe.dev.dtaclass.id/api/events/me', {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data;
  } catch (err: any) {
    throw new Error('Failed to fetch my events');
  }
};
