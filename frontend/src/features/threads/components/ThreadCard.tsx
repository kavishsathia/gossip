import {
  Card,
  CardActions,
  CardContent,
  CardMedia,
  IconButton,
} from "@mui/material";
import { Thread } from "../../../services/threads/types";
import { Heart, MessageCircle, Share2 } from "lucide-react";
import { Link, useSearchParams } from "react-router";

export default function ThreadCard({ item }: { item: Thread }) {
  const [searchParams] = useSearchParams();

  return (
    <Link
      className="w-full"
      to={`/thread/${item.ID}?${searchParams.toString()}`}
    >
      <Card
        className="w-full border border-gray-150 hover:shadow-lg transition-shadow duration-200"
        elevation={0}
      >
        <div className="flex gap-2">
          <CardMedia
            component="img"
            id={`${item.ID}-image`}
            sx={{ width: 140 }}
            image={item.Image || "https://placehold.co/400"}
          />

          <div className="flex flex-col flex-grow">
            <CardContent>
              <h3 className="text-base font-semibold mb-1 line-clamp-1 text-left">
                {item.Title}
              </h3>

              <h5 className="text-sm text-gray-800 line-clamp-2 text-left">
                {item.Body}
              </h5>
            </CardContent>

            <CardActions className="p-0 flex justify-between items-center">
              <div className="flex gap-4">
                <div className="flex items-center gap-1">
                  <IconButton className="p-1">
                    <Heart
                      color={item.Liked ? "red" : "black"}
                      className={`inline ${
                        item.Liked ? "fill-red-500" : "hover:fill-red-300"
                      } hover:scale-110 `}
                    />
                  </IconButton>
                  <span className="text-sm text-gray-600">{item.Likes}</span>
                </div>

                <div className="flex items-center gap-1">
                  <IconButton className="p-1">
                    <MessageCircle className="w-5 h-5 text-gray-600" />
                  </IconButton>
                  <span className="text-sm text-gray-600">{item.Comments}</span>
                </div>

                <div className="flex items-center gap-1">
                  <IconButton className="p-1">
                    <Share2 className="w-5 h-5 text-gray-600" />
                  </IconButton>
                  <span className="text-sm text-gray-600">{item.Shares}</span>
                </div>
              </div>
            </CardActions>
          </div>
        </div>
      </Card>
    </Link>
  );
}
