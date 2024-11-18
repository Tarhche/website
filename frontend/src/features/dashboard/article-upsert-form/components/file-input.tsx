import Image from "next/image";
import {useState} from "react";
import {Box, Stack, Paper, Modal, InputLabel, Button} from "@mantine/core";
import {FilesExplorer} from "@/components/files-explorer";
import {FILES_PUBLIC_URL} from "@/constants/envs";
import classes from "./file-input.module.css";

type Props = {
  name: string;
  label: string;
  defaultValue?: string;
  icon: React.ReactNode;
};

export function FileInput({name, label, defaultValue, icon}: Props) {
  const [isFileExplorerOpen, setIsFileExplorerOpen] = useState(false);
  const [selectedFile, setSelectedFile] = useState(defaultValue || "");

  const handleOpenFileExplorer = () => {
    setIsFileExplorerOpen(true);
  };

  const handleSelectFile = (fileId: string | undefined) => {
    setIsFileExplorerOpen(false);
    setSelectedFile(fileId || "");
  };

  const handleCloseFileExplorer = () => {
    setIsFileExplorerOpen(false);
  };

  return (
    <Box>
      <Stack gap={5}>
        <InputLabel>{label}</InputLabel>
        {!Boolean(selectedFile) ? (
          <Button
            classNames={{
              root: classes.fileInput,
            }}
            variant="default"
            onClick={handleOpenFileExplorer}
          >
            {icon}
          </Button>
        ) : (
          <Image
            src={`${FILES_PUBLIC_URL}/${selectedFile}`}
            alt="article's file"
            width={200}
            height={150}
            className={classes.image}
            onClick={handleOpenFileExplorer}
          />
        )}
      </Stack>
      <input name={name} value={selectedFile} hidden readOnly />
      <Modal
        size="xl"
        opened={isFileExplorerOpen}
        onClose={handleCloseFileExplorer}
        withCloseButton={false}
      >
        <Paper>
          <FilesExplorer onSelect={handleSelectFile} />
        </Paper>
      </Modal>
    </Box>
  );
}
