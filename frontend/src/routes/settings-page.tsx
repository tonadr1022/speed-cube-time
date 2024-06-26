import FileExportMenu from "../components/settings/FileExportMenu";
import SettingRow from "../components/settings/SettingRow";
import ModuleCountSelect from "../components/settings/ModuleCountSelect";
import { useState } from "react";
import ThemeMenu from "../components/settings/ThemeMenu";
import PageWrapper from "../components/layout/PageWrapper";
const options = ["Timer", "Appearance", "Data"];
const SettingsPage = () => {
  const [menuOption, setMenuOption] = useState("Timer");

  return (
    <PageWrapper title="Settings">
      <div className="max-w-md">
        <div className="mt-3">
          {options.map((option) => (
            <input
              key={option}
              type="radio"
              name="options"
              data-title={option}
              className="mr-2 btn btn-sm"
              value={option}
              aria-label={option}
              checked={menuOption === option}
              onChange={(e) => setMenuOption(e.target.value)}
            />
          ))}
        </div>
        <div className="flex flex-col">
          {menuOption === "Appearance" && (
            <SettingRow
              title="Theme"
              description="Change the theme of the app"
              flexRow={false}
            >
              <ThemeMenu />
            </SettingRow>
          )}
          {menuOption === "Timer" && (
            <SettingRow
              title="Number of Modules"
              description="Change the number of modules displayed on the timer screen"
            >
              <ModuleCountSelect />
            </SettingRow>
          )}
          {menuOption === "Data" && (
            <SettingRow
              title="Export Solves"
              description="Export all of your solves to CSV, TXT, or JSON"
            >
              <FileExportMenu />
            </SettingRow>
          )}
        </div>
      </div>
    </PageWrapper>
  );
};

export default SettingsPage;
