"use client";

import { useEffect, useState } from "react";
import FileInput from "./FileInput";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
  SelectGroup,
  SelectLabel,
} from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import MultipleSelector, { Option } from "../multiple-selector";
import "@styles/form.css";
import {Playbook, Setup, Rule, tradeService, formatOptionsToRules, formatRulesToOptions} from "@api/addTradeService"
import { useUser } from "@/context/userContext";
const OPTIONS: Option[] = [
  { label: "Break previous high", value: "idididid123" },
  { label: "Break current high", value: "ididid321" },
];



export default function AddTradeForm() {
  const [playbooks, setPlaybooks] = useState<Playbook[]>([]);
  const [setups, setSetups] = useState<Setup[]>([]);
  const [rules, setRules] = useState<Option[]>([]); //We will change rule to option later
  const [selectedPlaybook, setSelectedPlaybook] = useState<Playbook | null>(null);
  const [selectedSetup, setSelectedSetup] = useState<Setup | null>(null);
  const [selectedRules, setSelectedRules] = useState<Option[]>([]);
  const [loadingPlaybooks, setLoadingPlaybooks] = useState(false);
  const [loadingSetups, setLoadingSetups] = useState(false);
  const [loadingRules, setLoadingRules] = useState(false);
  const {user} = useUser();

  useEffect(() => {
    //TO DO: fetch from backend
    
    const getPlaybook = async () => {
      if(!user?.userId) return;
      setLoadingPlaybooks(true);
      try {
        const data = await tradeService.fetchPlaybooks(user.userId);
        setPlaybooks(data);
      } catch (err) {
        console.error("Failed to fetch playbooks", err);
      } finally {
        setLoadingPlaybooks(false);
      }
    }

    const getSetups = async () => {
      if(!selectedPlaybook) return;
      setLoadingSetups(true);
      try {
        const data = await tradeService.fetchSetups(selectedPlaybook.id);
        setSetups(data);
      } catch (err) {
        console.error("Failed to fetch setups", err);
      }
      finally {
        setLoadingSetups(false);
      }
    }
    const getRules = async () => {
      if(!selectedSetup) return;
      setLoadingRules(true);
      try {
        const ruleData = await tradeService.fetchRules(selectedSetup.id);
        const formatted : Option[] = formatRulesToOptions(ruleData);
        setRules(formatted);
      } catch (err) {
        console.error("Fail to fetch rules", err);
      } finally {
        setLoadingRules(false);
      }
    }
    getPlaybook();
    getSetups();
    getRules();
  }, [selectedSetup]);

  return (
    <form className="min-h-screen space-y-4 w-full bg-white p-5 rounded-lg flex flex-col items-center">
      <div className="w-full flex flex-col space-y-2">
        <Label className="self-start text-content-main text-base block pt-4">
          Add trade picture:
        </Label>

        <div className="w-full flex justify-center">
          <FileInput />
        </div>
      </div>

      <div className="input-line">
        <Label htmlFor="symbol" className="input-label">
          Symbol:
        </Label>
        <Input
          id="symbol"
          type="text"
          placeholder="e.g. XAU/USD US100"
          className="input-block"
        />
      </div>

      <div className="input-line">
        <Label htmlFor="entry_price" className="input-label">
          Entry price:
        </Label>
        <Input
          id="entry_price"
          type="number"
          step="0.01"
          placeholder="Entry Price"
          className="input-block"
        />
      </div>

      <div className="input-line">
        <Label htmlFor="symbol" className="input-label">
          Exit price:
        </Label>
        <Input
          id="exit_price"
          type="number"
          step="0.01"
          placeholder="Exit Price"
          className="input-block"
        />
      </div>

      <div className="input-line">
        <Label htmlFor="playbook" className="input-label">
          Playbook:
        </Label>
        <Select>
          <SelectTrigger>
            <SelectValue placeholder="Select a playbook" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel>Playbooks</SelectLabel>
              {playbooks.map((p) => (
                <SelectItem key={p.id} value={p.name}>
                  {" "}
                  {p.name}{" "}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <div className="input-line">
        <Label htmlFor="setup" className="input-label">
          Setup:
        </Label>
        <Select>
          <SelectTrigger >
            <SelectValue placeholder="Select a setup" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel>Setups</SelectLabel>
              {setups.map((s) => (
                <SelectItem key={s.id} value={s.name}>
                  {" "}
                  {s.name}{" "}
                </SelectItem>
              ))}
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <div className="input-line">
        <Label className="input-label">Criteria:</Label>
        <MultipleSelector
          className="w-full px-3"
          options={rules}
          placeholder="Select criteria met..."
          emptyIndicator={
            <p className="text-center text-md leading-10 text-gray-600 dark:text-gray-400">
              no results found.
            </p>
          }
        />
      </div>

      <button
        type="submit"
        className="bg-primary hover:bg-secondary px-4 py-2 rounded-lg"
      >
        Submit
      </button>
    </form>
  );
}
