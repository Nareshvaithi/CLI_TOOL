/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/shirou/gopsutil/v4/mem"
	"github.com/spf13/cobra"
)
var (
	critical float64
	criticalErr error
	warning float64
	warningErr error
)

// cpuUsageCmd represents the cpuUsage command
var cpuUsageCmd = &cobra.Command{
	Use:   "cpuUsage",
	Short: "A brief description of your command",
	Long: "TO FIND STATUS OF THE CPU command: cpuUsage --critical percent -- warning percent",
	Run: func(cmd *cobra.Command, args []string) {
		critical,criticalErr = cmd.Flags().GetFloat64("critical")
		warning,warningErr = cmd.Flags().GetFloat64("warning")
		if criticalErr != nil{
			fmt.Println(criticalErr)
			return
		}
		if warningErr != nil{
			fmt.Println(warningErr)
			return
		}
		fmt.Println(critical)
		fmt.Println(warning)
		cpuUsedPercent,cpuErr := findCpuUsage()
		if cpuErr != nil {
			log.Fatal(cpuErr)
		}
		status,exitCode := findStatus(cpuUsedPercent,critical,warning)
		fmt.Printf("%s[C:%.1f,W:%.1f]:CPU usage - %.2f%%\n",status,critical,warning,cpuUsedPercent)
		os.Exit(exitCode)
	},
}

func init() {
	rootCmd.AddCommand(cpuUsageCmd)
	cpuUsageCmd.Flags().Float64P("critical","c",90.0,"Critical Value of Percentage")
	cpuUsageCmd.Flags().Float64P("warning","w",80.0,"Warning Value of Percentage")
}

func findCpuUsage() (float64,error) {
	cpuData,err:= mem.VirtualMemory()
	Percent := cpuData.UsedPercent
	return Percent,err
}

func findStatus(cpuUsagePercent,critical,warning float64) (string,int) {
	var status string
	var exitCode int
	if cpuUsagePercent >= critical {
		status = "CRITICAL"
		exitCode = 2 
	}else if cpuUsagePercent >= warning {
		status = "WARNING"
		exitCode = 1
	}else {
		status = "OK"
		exitCode = 0
	}
	return status,exitCode
}