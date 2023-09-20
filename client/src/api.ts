import {mande} from "mande";

export const api = mande(window.location.href.includes('localhost') ? 'http://localhost:3000/api' : '/api');
