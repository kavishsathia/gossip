import "@mdxeditor/editor/style.css";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  TextField,
} from "@mui/material";
import { Upload } from "lucide-react";
import { useState } from "react";

const ImageSelectionModal = ({
  image,
  setImage,
}: {
  image: string;
  setImage: React.Dispatch<React.SetStateAction<string>>;
}) => {
  const [imageModal, setImageModal] = useState(false);
  const [imageBuffer, setImageBuffer] = useState("");

  return (
    <>
      <div
        onClick={() => setImageModal(true)}
        className="relative w-24 h-24 rounded-md cursor-pointer"
      >
        <img
          className="w-full h-full object-cover rounded-md"
          src={image}
          alt="Thread"
        />
        <span className="absolute top-0 left-0 w-full h-full flex items-center justify-center bg-gray-500/50 rounded-md text-white">
          <Upload className="stroke-2" />
        </span>
      </div>

      <Dialog open={imageModal} onClose={() => setImageModal(false)}>
        <DialogTitle>Add an Image</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Find an image on the Internet and add it here!
          </DialogContentText>
          <TextField
            value={imageBuffer}
            onChange={(e) => setImageBuffer(e.target.value)}
            autoFocus
            required
            margin="dense"
            label="Image URL"
            type="url"
            fullWidth
            variant="standard"
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setImageModal(false)}>Cancel</Button>
          <Button
            onClick={() => {
              setImage(imageBuffer);
              setImageBuffer("");
              setImageModal(false);
            }}
          >
            Add
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default ImageSelectionModal;
