import config from './config';
import pupeteer from "puppeteer"
import CSVParser from './utils/csv';

async function main() {
    try {
        const csvParser = new CSVParser("./data/example.csv");
        const { data } = await csvParser.parse();

        const browser = await pupeteer.launch({ headless: false, defaultViewport: { width: 1920, height: 1080 } });
        const page = await browser.newPage()
        await page.goto("https://autogestion.camaraargentina.com.ar/", { waitUntil: "networkidle2" })
        await page.click("body > div > div > div:nth-child(7) > div:nth-child(1) > a")
        await page.waitForSelector("body > div > div > div > div.col-md-6 > form > button")
        await page.$eval("#exampleInput1", (el: any) => el.value = "insaemoron2@gmail.com")
        await page.$eval("#exampleInputPassword1", (el: any) => el.value = "Perogrullo")
        await page.click("body > div > div > div > div.col-md-6 > form > button")

        // acta
        for (let i = 0; i < data.length; i++) {
            const row = data[i]
            if (i == 0) {
                await page.goto("https://autogestion.camaraargentina.com.ar/misactas")
                await page.waitForSelector("body > div.container-fluid > div:nth-child(2) > div > a")
                await page.click("body > div.container-fluid > div:nth-child(2) > div > a")
                await page.waitForSelector("#fisico")
                await page.$eval("#fisico", (el: any) => el.value = 0)
                await page.waitForSelector("#boton")
                await page.click("#boton")
                await page.waitForSelector("body > div.container-fluid > div:nth-child(2) > div:nth-child(2) > a")
                await page.click("body > div.container-fluid > div:nth-child(2) > div:nth-child(2) > a")
            }

            if (i > data.length - 1) {
                await page.goto("https://autogestion.camaraargentina.com.ar/misactas")
            }

            await page.waitForSelector("#form")
            await page.$eval("#form > div:nth-child(1) > div:nth-child(4) > div > input", (el: any, row: any) => el.value = row.name, row)
            await page.$eval("#form > div:nth-child(1) > div:nth-child(6) > div.col-md-6 > select", (el: any) => el.value = "DNI")
            await page.$eval("#form > div:nth-child(1) > div:nth-child(6) > div.col-md-3 > input", (el: any, row: any) => el.value = row.dni, row)
            await page.$eval("#form > div:nth-child(1) > div:nth-child(10) > div > input", (el: any, row: any) => el.value = row.phone, row)
            await page.$eval("#form > div:nth-child(1) > div:nth-child(12) > div > input", (el: any, row: any) => el.value = row.email, row)
            await page.$eval("#form > div:nth-child(1) > div:nth-child(14) > div > select", (el: any, row: any) => el.value = row.course, row)
            await page.$eval("#form > div:nth-child(1) > div:nth-child(16) > div:nth-child(2) > input", (el: any, row: any) => el.value = row.endDate, row)
            await page.$eval("#form > div:nth-child(1) > div:nth-child(16) > div:nth-child(4) > input", (el: any, row: any) => el.value = row.score, row)
            await page.waitForSelector("#boton")
            await page.click("#boton")

            await page.waitForSelector("body > div.container-fluid > div:nth-child(2) > div:nth-child(2) > a")
            await page.click("body > div.container-fluid > div:nth-child(2) > div:nth-child(2) > a")
        }

    }
    catch (error) {
        console.log(error);
        throw error;
    }
}

main()