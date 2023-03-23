import { Pagination, Typography } from "@mui/material";
import { Container } from "@mui/system";
import { Suspense } from "react";
import ErrorBoundary from "../Components/ErrorBoundary";
import { CountResource } from "../services/videos";

interface PaginationCounterProps {
    countResource: CountResource;
    page: number;
    setPage: (x: number) => void;
}

export default function PaginationCounter({ countResource, page, setPage }: PaginationCounterProps) {
    return (
        <Container sx={{ display: "flex" }}>
            <ErrorBoundary fallback={"error loading count"}>
                <Suspense fallback={"loading..."}>
                    <Typography
                        variant='subtitle1'>
                        total {countResource.read()} videos found
                    </Typography>
                    <Pagination
                        count={Math.ceil(countResource.read() / 20)}
                        page={page}
                        onChange={(_, pageNumber) => setPage(pageNumber)}
                        shape="rounded" />
                </Suspense>
            </ErrorBoundary>
        </Container>
    )
}