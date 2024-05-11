import { useState, useEffect } from "react";
import dayjs, { Dayjs } from "dayjs";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select from "@mui/material/Select";
import TopBar from "./TopBar";
import Grid from "@mui/material/Grid";
import Alert from "@mui/material/Alert";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import Chain from "./chain";
import { OptionChain } from "./types";
import { Paper } from "@mui/material";

function Grey() {
  const [alertOpen, setAlertOpen] = useState<Boolean>(false);
  const [alertMessage, setAlertMessage] = useState<string>("");
  const [symbol, setSymbol] = useState<string>("");
  const [date, setDate] = useState<Dayjs>(dayjs("2021-01-01"));
  const [expiry, setExpiry] = useState<string>("");
  const [availableExpiries, setExpiries] = useState<string[]>([]);
  const [chain, setChain] = useState<OptionChain[]>([]);

  useEffect(() => {
    if (symbol && date) {
      fetch(
        `https://grey.rebelscode.online/expiries/${symbol}/${date.format(
          "YYYY-MM-DD"
        )}`
      )
        .then(async (response) => {
          if (response.status !== 200) {
            console.log("error", symbol, date, response.status);
            setAlertOpen(true);
            setAlertMessage("Something Went Wrong!!!");
            return;
          }
          console.log("success", symbol, date);
          setAlertOpen(false);
          const data = await response.json();
          setExpiries(data.dates);
        })
        .catch((error) => {
          console.log("error", symbol, date, error);
          setAlertOpen(true);
          setAlertMessage("Something Went Wrong!!!");
          return;
        });
    }
  }, [symbol, date]);

  useEffect(() => {
    if (symbol && expiry && date) {
      fetch(
        `https://grey.rebelscode.online/chain/${symbol}/${date.format(
          "YYYY-MM-DD"
        )}/${expiry}`
      )
        .then(async (response) => {
          if (response.status !== 200) {
            console.log("error", symbol, expiry, date, response.status);
            setAlertOpen(true);
            setAlertMessage("Something Went Wrong!!!");
            return;
          }
          console.log("success", symbol, expiry, date);
          setAlertOpen(false);
          const data = await response.json();
          setChain(data);
        })
        .catch((error) => {
          console.log("error", symbol, expiry, date, error);
          setAlertOpen(true);
          setAlertMessage("Something Went Wrong!!!");
          return;
        });
    }
  }, [expiry, date]);

  return (
    <div>
      <TopBar />
      {alertOpen && <Alert severity="error">{alertMessage}</Alert>}
      <Container maxWidth="xl" sx={{ mt: 2 }}>
        <Grid container alignItems="center">
          <Grid item xs={12} md={4} sx={{ mt: 1 }}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <LocalizationProvider dateAdapter={AdapterDayjs}>
                <DatePicker
                  label="Start Date"
                  value={date}
                  onChange={(userDate) => userDate && setDate(userDate)}
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
                  onChange={(event) => setSymbol(event.target.value)}
                  label="Symbol"
                >
                  <MenuItem key={"NIFTY"} value={"NIFTY"}>
                    NIFTY
                  </MenuItem>
                  <MenuItem key={"BANKNIFTY"} value={"BANKNIFTY"}>
                    BANKNIFTY
                  </MenuItem>
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
                  onChange={(event) => setExpiry(event.target.value)}
                  label="Expiry"
                >
                  {availableExpiries.map(function (item) {
                    return (
                      <MenuItem key={item} value={item}>
                        {item}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Box>
          </Grid>
        </Grid>
      </Container>
      <Container maxWidth="xl" sx={{ mt: 1 }}>
        <Paper elevation={2}>
          <Grid container alignItems="center">
            <Grid item xs={12} md={4} sx={{ mt: 1 }}>
              <Box sx={{ display: "flex", justifyContent: "center" }}>
                <p>
                  {symbol}: {chain[1]?.spot}
                </p>
              </Box>
            </Grid>
            <Grid item xs={12} md={4} sx={{ mt: 1 }}>
              <Box sx={{ display: "flex", justifyContent: "center" }}>
                <p>Futures: {chain[1]?.fut}</p>
              </Box>
            </Grid>
            <Grid item xs={12} md={4} sx={{ mt: 1 }}>
              <Box sx={{ display: "flex", justifyContent: "center" }}>
                <p>VIX: {chain[1]?.vix}</p>
              </Box>
            </Grid>
          </Grid>
        </Paper>
      </Container>
      <Container maxWidth="xl" sx={{ mt: 1 }}>
        <Paper elevation={2}></Paper>
      </Container>
      <Container maxWidth="xl" sx={{ mt: 1 }}>
        <Chain chain={chain} />
      </Container>
    </div>
  );
}

export default Grey;
