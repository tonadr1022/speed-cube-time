import React from "react";

type Props = {
  title: string;
  children: React.ReactNode;
};

const PageWrapper = ({ title, children }: Props) => {
  return (
    <div className="h-full p-6 overflow-y-auto">
      <h1 className="text-6xl font-semibold">{title}</h1>
      {children}
    </div>
  );
};

export default PageWrapper;
