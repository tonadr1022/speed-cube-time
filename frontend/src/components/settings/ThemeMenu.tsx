import { useSettings } from "../../hooks/useContext";
import clsx from "clsx";
const options = [
  "light",
  "dark",
  "cupcake",
  "bumblebee",
  "emerald",
  "corporate",
  "synthwave",
  "retro",
  "cyberpunk",
  "valentine",
  "halloween",
  "garden",
  "forest",
  "aqua",
  "lofi",
  "pastel",
  "fantasy",
  "wireframe",
  "black",
  "luxury",
  "dracula",
  "autumn",
  "business",
  "acid",
  "lemonade",
  "night",
  "winter",
];

const ThemeMenu = () => {
  const { theme, setTheme } = useSettings();
  const handleClick = (option: string) => {
    setTheme(option);
  };
  return (
    <div className="grid grid-cols-4 ">
      {options.map((option) => (
        <button
          key={option}
          // data-set-theme={option}
          // data-act-class={theme === "ACTIVECLASS"}
          className={clsx(
            "bg-base-300 rounded-lg p-2 text-center text-xs font-semibold cursor-pointer m-1",
            { "ring-2 ring-primary": theme === option },
          )}
          onClick={() => handleClick(option)}
        >
          {option}
        </button>
      ))}
    </div>
  );
};

export default ThemeMenu;
