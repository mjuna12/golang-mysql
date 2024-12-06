import { defineStore } from "pinia";
import Cookies from "js-cookie";
import { EventResponseTypes } from "../types/auth.type";
import {
  registerApi,
  loginApi,
  getEventListApi,
  registerEventApi,
  getMyEventApi,
} from "../api/auth_api";

// Define the Pinia store
export const useAuth = defineStore("auth", {
  state: () => ({
    token: "",
    error: "" as any,
    list_event: {} as EventResponseTypes,
    list_myevent: {} as any,
    selectedEvent : {} as any,
  }),

  getters: {
    getToken: (state) => {
      const tokenCookie = Cookies.get("token");
      return state.token || (tokenCookie ? JSON.parse(tokenCookie) : null);
    },
    getIsLogin: (state) => {
      const tokenCookie = Cookies.get("token");
      return !!state.token || !!tokenCookie;
    },
  },

  actions: {
    setToken(token: string) {
      this.token = token;
      Cookies.set("token", JSON.stringify(this.token), { expires: 1 });
    },

    // Register action
    async register(email: string, password: string) {
      this.error = null;
      try {
        await registerApi(email, password);
        return true;
      } catch (err: any) {
        this.error = err.message;
        return false;
      }
    },

    // Login action
    async login(email: string, password: string) {
      this.error = null;
      try {
        const response = await loginApi(email, password);
        this.setToken(response.token);
        return true;
      } catch (err: any) {
        this.error = err.message;
        return false;
      }
    },

    // Logout action
    signOut() {
      this.token = "";
      Cookies.remove("token");
    },

    // Get Event List action
    async getEventList() {
      const token = this.getToken;
      if (token) {
        const data = await getEventListApi(token);
        if (data?.data !== null) {
          this.list_event = data;
        }
      } else {
        this.error = { general: ["Token not found"] };
      }
    },

    // Register Event action
    async registerEvent(eventId: string, lat: number, lng: number) {
      try {
        const response = await registerEventApi(
          eventId,
          lat,
          lng,
          this.getToken
        );
        return response;
      } catch (err: any) {
        this.error = err.message;
        return null;
      }
    },
    setSelectedEvent(value:any) {
      this.selectedEvent = value
    },
    // Get My Events action
    async getMyEvent() {
      try {
        const token = this.getToken;
        if (token) {
          const data = await getMyEventApi(token);
          if (data?.data !== null) {
            this.list_myevent = data;
          }
        } else {
          this.error = { general: ["Token not found"] };
        }
      } catch (err: any) {
        this.error = err.message;
        return null;
      }
    },
  },
});
