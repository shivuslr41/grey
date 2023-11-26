import { useState } from "react";
import dayjs, { Dayjs } from "dayjs";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import TopBar from "./TopBar";
import Grid from "@mui/material/Grid";
import Alert from "@mui/material/Alert";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import Chain from "./chain";

type OptionChain = {
  cgamma: number
  cvega: number
  ctheta: number
  cdelta: number
  ce: number
  pe: number
  pdelta: number
  ptheta: number
  pvega: number
  pgamma: number
  strike: number
  spot: number
  fut: number
  vix: number
}

function Grey() {
  const [alertOpen, setAlertOpen] = useState<Boolean>(false);
  const [alertMessage, setAlertMessage] = useState<string>("");
  const [symbol, setSymbol] = useState<string>("");
  const [date, setDate] = useState<Dayjs>(dayjs("2021-01-01"));
  const [expiry, setExpiry] = useState<string>("");
  const [availableExpiries, setExpiries] = useState<string[]>([]);
  const [chain, setChain] = useState<OptionChain[]>([]);

  const getExpiryList = (event: SelectChangeEvent) => {
    setSymbol(event.target.value);
    let symbol = event.target.value;
    fetch(
      `https://grey.rebelscode.online/expiries/${symbol}/${date.format(
        "YYYY-MM-DD",
      )}`,
    )
      .then(async (response) => {
        if (response.status !== 200) {
          console.log("error", symbol, response.status);
          setAlertOpen(true);
          setAlertMessage("Something Went Wrong!!!");
          return;
        }
        console.log("success", symbol);
        setAlertOpen(false);
        const data = await response.json();
        setExpiries(data.dates);
      })
      .catch((error) => {
        console.log("error", symbol, error);
        setAlertOpen(true);
        setAlertMessage("Something Went Wrong!!!");
        return;
      });
  };

  const getOptionChain = (event: SelectChangeEvent) => {
    setExpiry(event.target.value);
    let expiry = event.target.value;
    fetch(
      `https://grey.rebelscode.online/expiries/${symbol}/${date.format(
        "YYYY-MM-DD",
      )}`,
    )
      .then(async (response) => {
        if (response.status !== 200) {
          console.log(
            "error",
            symbol,
            expiry,
            response.status,
            response.text(),
          );
          setAlertOpen(true);
          setAlertMessage("Something Went Wrong!!!");
          return;
        }
        console.log("success", symbol, expiry);
        setAlertOpen(false);
        const data = await response.json();
        setExpiries(data.dates);
      })
      .catch((error) => {
        console.log("error", symbol, expiry, error);
        setAlertOpen(true);
        setAlertMessage("Something Went Wrong!!!");
        return;
      });
    console.log(symbol, expiry, event.target.value);
  };

  return (
    <div>
      <TopBar />
      {alertOpen && <Alert severity="error">{alertMessage}</Alert>}
      <Container maxWidth="xl" sx={{ mt: 2 }}>
        <Grid container alignItems="center" justify="center">
          <Grid item xs={12} md={4} sx={{ mt: 1 }}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <LocalizationProvider dateAdapter={AdapterDayjs}>
                <DatePicker
                  label="Start Date"
                  value={date}
                  onChange={(userDate) => setDate(userDate)}
                />
              </LocalizationProvider>
            </Box>
          </Grid>
          <Grid item xs={12} md={4}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
                <InputLabel id="symbol-label">Symbol</InputLabel>
                <Select
                  labelId="symbol-label"
                  value={symbol}
                  onChange={getExpiryList}
                  label="Symbol"
                >
                  <MenuItem value={"NIFTY"}>NIFTY</MenuItem>
                  <MenuItem value={"BANKNIFTY"}>BANKNIFTY</MenuItem>
                </Select>
              </FormControl>
            </Box>
          </Grid>
          <Grid item xs={12} md={4}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
                <InputLabel id="expiry-label">Expiry</InputLabel>
                <Select
                  labelId="expiry-label"
                  value={expiry}
                  onChange={getOptionChain}
                  label="Expiry"
                >
                  {availableExpiries.map(function (item) {
                    return <MenuItem value={item}>{item}</MenuItem>;
                  })}
                </Select>
              </FormControl>
            </Box>
          </Grid>
        </Grid>
      </Container>
      <Container maxWidth="xl" sx={{ mt: 1 }}>
        <Chain />
      </Container>
    </div>
  );
}

export default Grey;
