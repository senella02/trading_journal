import React, { useRef, useState } from "react";
import { IoIosAddCircle, IoMdCloudUpload } from "react-icons/io";
import { MdDelete } from "react-icons/md";
export default function FileInput() {
  const inputRef = useRef<HTMLInputElement | null>(null);

  const [selectedFile, setSelectedFile] = useState<File | null>(null);

  interface FileInputChangeEvent extends React.ChangeEvent<HTMLInputElement> {}

  const handleOnChange = (e: FileInputChangeEvent): void => {
    if (e.target.files && e.target.files.length > 0) {
      console.log("Selected file:", e.target.files[0].name); // Log file name
      setSelectedFile(e.target.files[0]);
    }
  };

  const onChooseFile = () => {
    console.log("File button clicked");
    inputRef.current?.click();
  };

  const removeFile = () => {
    setSelectedFile(null);
    if (inputRef.current) {
      inputRef.current.value = ""; // Reset input value
    }
  };

  return (
    <div className="flex flex-col justify-center items-center gap-2 p-4 max-w-full">
      <input
        type="file"
        ref={inputRef}
        onChange={handleOnChange}
        className="hidden"
        accept="*"
      ></input>
      
        <button
          type="button"
          onClick={onChooseFile}
          className="flex flex-col items-center justify-center gap-2 rounded-3xl w-full h-[180px] bg-white border-black border-[3px] border-dashed"
        >
          <IoMdCloudUpload className="text-4xl text-content-main" />
          <p className="text-xl text-content-main">Upload file</p>
        </button>
      
      
      {selectedFile && (
        <div className="inline-flex flex-1 items-center gap-2 mt-2 bg-gray-100 p-2 rounded max-w-full">
          <p className="text-black truncate max-w-[300px]">{selectedFile.name}</p>
          <button onClick={removeFile}>
            <MdDelete color="black" />
          </button>
        </div>
      )}
    </div>
  );
}
