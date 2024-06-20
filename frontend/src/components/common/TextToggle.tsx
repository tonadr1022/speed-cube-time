import clsx from "clsx";

type Props = {
  title1: string;
  title2: string;
  name: string;
  on: boolean;
  className?: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const TextToggle = ({
  name,
  title1,
  title2,
  on,
  onChange,
  className,
}: Props) => {
  return (
    <div className={clsx("join text-center my-1", className)}>
      <input
        type="radio"
        name={name}
        data-title={title1}
        className="join-item btn btn-xs"
        value={title1}
        aria-label={title1}
        checked={!on}
        onChange={onChange}
      />
      <input
        type="radio"
        name={name}
        data-title={title2}
        className="join-item btn btn-xs"
        value={title2}
        aria-label={title2}
        checked={on}
        onChange={onChange}
      />
    </div>
  );
};

export default TextToggle;
