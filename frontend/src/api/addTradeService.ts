import { Option } from "@/components/ui/multiple-selector";
import axios from "axios";
const API_BASE_URL = process.env.NEXT_PUBLIC_BACKEND_URL || "http://localhost:8000";
const UPLOAD_API_BASE_URL = "http://localhost:5433";
//Define type for data we want to send/get
export interface TradeData {
  symbol: string;
  value: string;
  entry_price: number;
  exit_price: number;
  playbook_id: string;
  setup_id: string;
  rules: string[];
  image_url?: string;
}

export interface Playbook {
  id: string;
  name: string;
}

export interface Setup {
  id: string;
  name: string;
}

export interface Rule {
  id: string;
  name: string;
}

export async function uplaodImage(file: File): Promise<string> {
  const formData = new FormData();
  formData.append("image", file);

  const response = await axios.post(`${UPLOAD_API_BASE_URL}/api/v1/images`, formData, {
    headers: { "Content-Type" : "multipart/form-data"},
  });

  return response.data.url;
}

export function formatOptionsToRules(options: Option[]) : Rule[] {
  const formatted : Rule[] = options.map((o) => ({
    id : String(o.id),
    name: o.label
  }));
  return formatted;
}

export function formatRulesToOptions(rules: Rule[]) : Option[] {
  const formatted: Option[] = rules.map((r) => ({
    label: r.name,
    value: String(r.id)
  }));
  return formatted;
}

export const tradeService = {
  fetchPlaybooks: async (userId : string): Promise<Playbook[]> => {
    const response = await axios.get(`${API_BASE_URL}/playbooks`);
    return response.data;
  },
  //fetch setup and rules after playbook
  fetchSetups: async (playbookId : string): Promise<Setup[]> => {
    const response = await axios.get(`${API_BASE_URL}/playbooks/${playbookId}/setups`);
    return response.data;
  },

  fetchRules: async (setupId: string): Promise<Rule[]> => {
    const response = await axios.get(`${API_BASE_URL}/setups/${setupId}/rules`);
    return response.data;
  },
  //If we have image, we will upload image first
  createTrade: async (image : File, data : TradeData) : Promise<any> => {
    let url : string = "";
    if(image){
      url = await uplaodImage(image); //Promise require await!
    }
    const formData = new FormData();
    formData.append("symbol", data.symbol);
    formData.append("entry_price", data.entry_price.toString());
    formData.append("exit_price", data.exit_price.toString());
    formData.append("playbook_id", data.playbook_id);
    formData.append("setup_id", data.setup_id);
    data.rules.forEach((c) => formData.append("rules", c));
    if(url != "") formData.append("url", url);
    const res = await axios.post(`${API_BASE_URL}/api/v1/trades`, data);
    return res.data;
  },
  uplaodImage
};