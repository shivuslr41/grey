import Accordion from "@mui/material/Accordion";
import AccordionSummary from "@mui/material/AccordionSummary";
import AccordionDetails from "@mui/material/AccordionDetails";
import Typography from "@mui/material/Typography";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { OptionChain } from "./types";

interface ChainProps {
  chain: OptionChain[];
}

export default function Chain(ChainProps: ChainProps) {
  return (
    <div>
      <Accordion>
        <AccordionSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel1a-content"
          id="panel1a-header"
        >
          <Typography>Option Chain</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <TableContainer component={Paper}>
            <Table
              sx={{ minWidth: 650 }}
              size="small"
              aria-label="a dense table"
            >
              <TableHead>
                <TableRow>
                  <TableCell>Gamma</TableCell>
                  <TableCell>Vega</TableCell>
                  <TableCell>Theta</TableCell>
                  <TableCell>Delta</TableCell>
                  <TableCell>CE</TableCell>
                  <TableCell>Strike</TableCell>
                  <TableCell>PE</TableCell>
                  <TableCell>Delta</TableCell>
                  <TableCell>Theta</TableCell>
                  <TableCell>Vega</TableCell>
                  <TableCell>Gamma</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {Array.isArray(ChainProps.chain) &&
                  ChainProps.chain.map(function (c: OptionChain, i) {
                    if (i == 0) {
                      return <></>;
                    }
                    return (
                      <TableRow
                        key={c.strike}
                        sx={{
                          "&:last-child td, &:last-child th": { border: 0 },
                        }}
                      >
                        <TableCell>{c.cgamma}</TableCell>
                        <TableCell>{c.cvega}</TableCell>
                        <TableCell>{c.ctheta}</TableCell>
                        <TableCell>{c.cdelta}</TableCell>
                        <TableCell>{c.ce}</TableCell>
                        <TableCell component="th" scope="row">
                          {c.strike}
                        </TableCell>
                        <TableCell>{c.pe}</TableCell>
                        <TableCell>{c.pdelta}</TableCell>
                        <TableCell>{c.ptheta}</TableCell>
                        <TableCell>{c.pvega}</TableCell>
                        <TableCell>{c.pgamma}</TableCell>
                      </TableRow>
                    );
                  })}
              </TableBody>
            </Table>
          </TableContainer>
        </AccordionDetails>
      </Accordion>
    </div>
  );
}
