import Link from "next/link";
import {Box, Stack, Alert, Anchor} from "@mantine/core";
import {AuthGuard} from "@/components/auth-guard";
import {CommentForm} from "./comment-form";
import {Comment} from "./comment";
import {IconInfoCircle} from "@tabler/icons-react";
import {fetchArticleComments} from "@/dal/comments";

type Props = {
  uuid: string;
};

export async function Comments({uuid}: Props) {
  const comments = (await fetchArticleComments(uuid)).items;
  const rootComments = comments.filter((c) => c.parent_uuid === undefined);

  return (
    <>
      <AuthGuard
        fallback={
          <Alert
            mt={"md"}
            variant="light"
            color="yellow"
            title="نیازمند احراز هویت"
            icon={<IconInfoCircle />}
          >
            برای اینکه بتوانید دیدگاه خود را ثبت کنید باید ابتدا{" "}
            <Anchor underline="always" href={"/auth/login"} component={Link}>
              وارد حسابتان
            </Anchor>{" "}
            شوید
          </Alert>
        }
      >
        <Box mt={"lg"}>
          <CommentForm object_uuid={uuid} parent_uuid="" />
        </Box>
      </AuthGuard>
      <Stack mt={"xl"}>
        {rootComments.map((comment) => {
          return (
            <Comment key={comment.uuid} comments={comments} comment={comment} />
          );
        })}
        {comments.length === 0 && (
          <Alert variant="light" color="green" icon={<IconInfoCircle />}>
            هنوز دیدگاهی برای این مقاله ثبت نشده!
          </Alert>
        )}
      </Stack>
    </>
  );
}