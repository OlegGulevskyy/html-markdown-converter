import { useState } from "react";
import {
  Run,
  FetchSpreadsheetInfo,
  FolderPicker,
} from "../../../wailsjs/go/main/App";
import { main } from "../../../wailsjs/go/models";
import { Input } from "../../components/input";

import { Spinner, WarningTriangle, Check } from "../../assets/svg";
import { Modal } from "../../components/modal";
import { SpinnerWhite } from "../../assets/svg/Spinner";

const testSheetName = "Templates";
const range = (a1Notation: string) => `${testSheetName}!${a1Notation}`;

const inputsDefState = {
  spreadsheetId: "1OCbB2tYqcWt_qzfGo43CfFAPOGvA_iJCq3aGq1hUaRA",
  htmlBodyRange: range("C2:C"),
  articleTitlesRange: range("B2:B"),
  categoryTitleRange: range("A2:A"),
  destinationPath: "/Users/oleggulevskyy/Desktop/articles-results",
  imagesPath: "/Users/oleggulevskyy/Desktop/articles-results/images/",
};

type Inputs = typeof inputsDefState;

export function Home() {
  const [inputs, setInputs] = useState(inputsDefState);
  const [ssTitle, setSsTitle] = useState<string | null>("");
  const [isLoadingSsInfo, setIsLoadingSsInfo] = useState(false);
  const [isTransforming, setIsTransforming] = useState(false);
  const [transformationDone, setTransformationDone] = useState(false);
  const [
    transformationData,
    setTransformationData,
  ] = useState<main.OperationRunStatus | null>(null);

  const resetAllFields = () => setInputs(inputsDefState);

  const runApp = (e: React.FormEvent<HTMLButtonElement>) => {
    e.preventDefault();
    setIsTransforming(true);
    Run({
      html_body_range: inputs.htmlBodyRange,
      spreadsheet_id: inputs.spreadsheetId,
      article_name_sheet_range: inputs.articleTitlesRange,
      md_dest_folder: inputs.destinationPath,
      category_name_sheet_range: inputs.categoryTitleRange,
      images_dest_folder: inputs.imagesPath,
    })
      .then((transformData) => {
        setTransformationData(transformData);
        setTransformationDone(true);
        setIsTransforming(false);
      })
      .catch((e) => console.log("Error transforming", e))
      .finally(() => setIsTransforming(false));
  };

  const onInput = (key: keyof Inputs, value: string) => {
    setInputs({
      ...inputs,
      [key]: value,
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

  const openFolderDialog = async (
    key: keyof Pick<Inputs, "destinationPath" | "imagesPath">
  ) => {
    const result = await FolderPicker();
    if (!result) return;

    onInput(key, result);
  };

  return (
    <div className="p-8 bg-slate-200">
      {!isTransforming && transformationDone && (
        <Modal transformationData={transformationData} />
      )}
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
                      onInput={(e) =>
                        onInput("spreadsheetId", e.currentTarget.value)
                      }
                      disabled={isTransforming}
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
                    onInput={(e) =>
                      onInput("htmlBodyRange", e.currentTarget.value)
                    }
                    label="HTML body range"
                    id="html-body-range"
                    disabled={isTransforming}
                  />
                </div>

                <div className="sm:col-span-3">
                  <Input
                    value={inputs.articleTitlesRange}
                    id="article-title-range"
                    onInput={(e) =>
                      onInput("articleTitlesRange", e.currentTarget.value)
                    }
                    label="Article name range"
                    disabled={isTransforming}
                  />
                </div>

                <div className="sm:col-span-3">
                  <Input
                    value={inputs.categoryTitleRange}
                    id="category-title-range"
                    onInput={(e) =>
                      onInput("categoryTitleRange", e.currentTarget.value)
                    }
                    label="Category name range"
                    disabled={isTransforming}
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
                        disabled={isTransforming}
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
                <div>
                  <label
                    htmlFor="email"
                    className="block text-sm font-medium text-gray-700 mb-2"
                  >
                    Destination project folder
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <button
                      type="button"
                      className="relative -ml-px inline-flex items-center space-x-2 rounded-l-md border border-gray-300 bg-gray-50 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
                      onClick={() => openFolderDialog("destinationPath")}
                      disabled={isTransforming}
                    >
                      <span>Select</span>
                    </button>
                    <div className="relative flex flex-grow items-stretch focus-within:z-10">
                      <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3"></div>
                      <input
                        type="email"
                        name="email"
                        id="email"
                        className="block w-full rounded-none rounded-r-md border-gray-300 pl-4 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                        placeholder="~/Desktop/my-project"
                        value={inputs.destinationPath}
                        onInput={(e) =>
                          onInput("destinationPath", e.currentTarget.value)
                        }
                        disabled={isTransforming}
                      />
                    </div>
                  </div>
                </div>

                <div className="mt-4">
                  <label
                    htmlFor="email"
                    className="block text-sm font-medium text-gray-700 mb-2"
                  >
                    Images folder
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <button
                      type="button"
                      className="relative -ml-px inline-flex items-center space-x-2 rounded-l-md border border-gray-300 bg-gray-50 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
                      onClick={() => openFolderDialog("imagesPath")}
                      disabled={isTransforming}
                    >
                      <span>Select</span>
                    </button>
                    <div className="relative flex flex-grow items-stretch focus-within:z-10">
                      <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3"></div>
                      <input
                        type="email"
                        name="email"
                        id="email"
                        className="block w-full rounded-none rounded-r-md border-gray-300 pl-4 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                        placeholder="~/Desktop/my-project/_images"
                        value={inputs.imagesPath}
                        onInput={(e) =>
                          onInput("imagesPath", e.currentTarget.value)
                        }
                        disabled={isTransforming}
                      />
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="pt-5">
            <div className="flex justify-between">
              <div>
                <p className="text-slate-400 text-sm">Made with ❤️ by Oleg</p>
                <p className="text-slate-400 text-sm">Powered by Wails</p>
              </div>
              <div className="">
                <button
                  type="button"
                  className="rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                  onClick={resetAllFields}
                  disabled={isTransforming}
                >
                  Reset all fields
                </button>
                <button
                  type="submit"
                  className={`ml-3 inline-flex justify-center rounded-md border border-transparent ${
                    !isTransforming ? "bg-indigo-600" : "bg-slate-500"
                  } py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2`}
                  onClick={runApp}
                  disabled={isTransforming}
                >
                  {isTransforming && (
                    <SpinnerWhite className="animate-spin h-4 w-4 text-white inline-flex mr-2" />
                  )}
                  Run transformation
                </button>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
}
