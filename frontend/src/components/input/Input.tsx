import { FC } from "react";

type InputProps = {
  onInput: (e: React.FormEvent<HTMLInputElement>) => any;
  label?: string;
  value: string;
	id: string;
	disabled?: boolean
};

export const Input: FC<InputProps> = ({ onInput, label, value, id, disabled }) => {
  return (
    <>
      {label && (
        <label
          htmlFor={id}
          className="block text-sm font-medium text-gray-700"
        >
          {label}
        </label>
      )}
      <div className="mt-1">
        <input
          onInput={onInput}
          value={value}
          type="text"
          name="html-body-range"
          id={id}
					disabled={disabled}
          autoComplete="given-name"
          className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        />
      </div>
    </>
  );
};
