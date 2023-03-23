import { Pagination, Typography, useTheme } from "@mui/material";
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
    const theme = useTheme()
    return (
        <Container sx={{
            display: "flex",
            padding: theme.spacing(1),
            alignItems: "center",
            justifyContent: "center"
        }}>
            <Typography variant="subtitle1"> {countResource.read()} </Typography>

            <Typography
                variant='subtitle1' sx={{ marginLeft : theme.spacing(1) }}>
                videos found
            </Typography>
            <Pagination
                count={Math.ceil(countResource.read() / 20)}
                page={page}
                onChange={(_, pageNumber) => setPage(pageNumber)}
                shape="rounded" />
        </Container>
    )
}