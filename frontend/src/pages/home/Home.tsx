import { useState } from "react";
import { Run, FetchSpreadsheetInfo } from "../../../wailsjs/go/main/App";
import { Input } from "../../components/input";

import { Spinner, WarningTriangle, Check } from "../../assets/svg";

const inputsDefState = {
  spreadsheetId: "1FFK1zo901uI_ZwuOLFCf7GSBl2UEcb6MvelR3Y0Oeh4",
  htmlBodyRange: "X2:X",
  articleTitlesRange: "N2:N",
  haveSharedSpreadsheet: "",
  destinationPath: "",
  imagesPath: "",
};

export function Home() {
  const [inputs, setInputs] = useState(inputsDefState);
  const [ssTitle, setSsTitle] = useState<string | null>("");
  const [isLoadingSsInfo, setIsLoadingSsInfo] = useState(false);

  const runApp = (e: React.FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    Run(inputs.spreadsheetId, inputs.htmlBodyRange, inputs.articleTitlesRange);
  };

  const onInput = (
    key: keyof typeof inputsDefState,
    e: React.FormEvent<HTMLInputElement>
  ) => {
    setInputs({
      ...inputs,
      [key]: e.currentTarget.value,
    });
  };

  const fetchSpreadsheetInfo = () => {
    if (!inputs.spreadsheetId) return;
    setSsTitle("");

    setIsLoadingSsInfo(true);
    FetchSpreadsheetInfo(inputs.spreadsheetId)
      .then((title) => setSsTitle(title))
      .catch(() => setSsTitle(null))
      .finally(() => setIsLoadingSsInfo(false));
  };

  return (
    <div className="p-8 bg-slate-200">
      <div className="bg-white p-6 rounded-md shadmw-md">
        <form className="space-y-8 divide-y divide-gray-200">
          <div className="space-y-8 divide-y divide-gray-200">
            <div>
              <div>
                <h3 className="text-lg font-medium leading-6 text-gray-900">
                  HTML to Markdown
                </h3>
                <p className="mt-1 text-sm text-gray-500">
                  Convert HTML to Markdown (.md or .mdx) files
                </p>
              </div>
            </div>

            <div className="pt-6">
              <div>
                <h3 className="text-lg font-medium leading-6 text-gray-900">
                  Data source
                </h3>
                <p className="mt-1 text-sm text-gray-500">
                  Spreadsheet information
                </p>
              </div>
              <div className="mt-6 grid grid-cols-1 gap-y-6 gap-x-4 sm:grid-cols-6">
                <div className="sm:col-span-6">
                  <label
                    htmlFor="username"
                    className="block text-sm font-medium text-gray-700"
                  >
                    Spreadsheet ID
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <span className="inline-flex items-center rounded-l-md border border-r-0 border-gray-300 bg-gray-50 px-3 text-gray-500 sm:text-sm">
                      https://docs.google.com/spreadsheets/d/
                    </span>
                    <input
                      type="text"
                      name="spreadsheet-id"
                      id="spreadsheet-id"
                      value={inputs.spreadsheetId}
                      autoComplete="spreadsheet-id"
                      onInput={(e) => onInput("spreadsheetId", e)}
                      onBlur={fetchSpreadsheetInfo}
                      className="block w-full min-w-0 flex-1 rounded-none rounded-r-md border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    />
                  </div>

                  <div className="mt-1">
                    <p className="text-slate-500 text-sm">
                      {isLoadingSsInfo && (
                        <>
                          <Spinner className="animate-spin mr-1 h-4 w-4 text-white inline" />
                          Loading...
                        </>
                      )}
                      {!isLoadingSsInfo && ssTitle === "" && (
                        <>I will check if Spreadsheet ID is accessible</>
                      )}
                      {!isLoadingSsInfo && ssTitle === null && (
                        <>
                          <WarningTriangle className="text-red-500 mr-1 h-4 w-4 inline" />
                          No spreadsheet with given ID found
                        </>
                      )}
                      {!isLoadingSsInfo && ssTitle && (
                        <>
                          <Check className="text-green-500 mr-1 h-4 w-4 inline" />
                          Found spreadsheet{" "}
                          <span className="font-bold">{ssTitle}</span>
                        </>
                      )}
                    </p>
                  </div>
                </div>
                <div className="sm:col-span-3">
                  <Input
                    value={inputs.htmlBodyRange}
                    onInput={(e) => onInput("htmlBodyRange", e)}
                    label="HTML body range"
                    id="html-body-range"
                  />
                </div>

                <div className="sm:col-span-3">
                  <Input
                    value={inputs.articleTitlesRange}
                    id="article-title-range"
                    onInput={(e) => onInput("articleTitlesRange", e)}
                    label="Article name range"
                  />
                </div>

                <div className="sm:col-span-6">
                  <div className="relative flex items-start">
                    <div className="flex h-5 items-center">
                      <input
                        id="comments"
                        name="comments"
                        type="checkbox"
                        className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                      />
                    </div>
                    <div className="ml-3 text-sm">
                      <label
                        htmlFor="comments"
                        className="font-medium text-gray-700"
                      >
                        I have shared access to the spreadsheet
                      </label>
                      <p className="text-gray-500">
                        For script to be able to access the data in your
                        spreadsheet, please add{" "}
                        <span className="font-bold">
                          html-to-markdown-extractor@dev-form-publisher.iam.gserviceaccount.com{" "}
                        </span>
                        as "Can view" to your spreadsheet.
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div className="pt-6">
              <div>
                <h3 className="text-lg font-medium leading-6 text-gray-900">
                  Destination
                </h3>
                <p className="mt-1 text-sm text-gray-500">
                  Define where generated files will be stored and how
                </p>
              </div>
              <div className="mt-6">
                <Input
                  value={inputs.destinationPath}
                  onInput={(e) => onInput("destinationPath", e)}
                  id="destination-path"
                  label="Project folder"
                />
              </div>
              <div className="mt-6">
                <Input
                  value={inputs.imagesPath}
                  onInput={(e) => onInput("imagesPath", e)}
                  id="image-path"
                  label="Images folder"
                />
              </div>
            </div>
          </div>

          <div className="pt-5">
            <div className="flex justify-end">
              <button
                type="button"
                className="rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
              >
                Reset all fields
              </button>
              <button
                type="submit"
                className="ml-3 inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                onClick={runApp}
              >
                Run transformation
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
}
