export type RegisterResponseTypes = {
  message: string;
  results: string;
};

export type LoginResponseTypes = {
  message: string;
  results: {
    access_token: string;
    token_type: string;
    expires_in: string;
  };
};

export type EventTypes = {
  event_id: number;
  event_name: string;
  event_picture: string;
  event_desc: string | any;
  event_is_close: boolean | null | any;
  event_start: string;
  event_end: string;
};

export type PaginationTypes = {
  total: number;
  count: number;
  per_page: number;
  current_page: number;
  prev_page: string | null;
  next_page: string | null;
  total_pages: number;
};

export type EventResponseTypes = {
  data: {
    message: string;
    results: Event[];
    pagination: PaginationTypes;
  };
};
