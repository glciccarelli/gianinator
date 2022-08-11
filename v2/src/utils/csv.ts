import fs from "fs"
import papa from "papaparse"

export default class CSVParser {
    private filePath: string

    constructor(filePath: string) {
        this.filePath = filePath;
    }

    public async parse(): Promise<papa.ParseResult<unknown>> {
        const csvFile = fs.readFileSync(this.filePath)
        const csvData = csvFile.toString()
        return papa.parse(csvData, {
            header: true,
            skipEmptyLines: true,
            complete: (results) => {
                console.log('Complete', results.data.length, 'records.');
                return results
            }
        })
    }
}