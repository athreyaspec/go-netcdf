package netcdf_test

import (
	"fmt"
	"log"

	"github.com/athreyaspec/go-netcdf/netcdf"
)

// CreateExampleFile creates an example NetCDF file containing only one variable.
func CreateExampleFile(filename string) error {
	// Create a new NetCDF 4 file. The dataset is returned.
	ds, err := netcdf.CreateFile("gopher.nc", netcdf.CLOBBER|netcdf.NETCDF4)
	if err != nil {
		return err
	}
	defer ds.Close()

	// Add the dimensions for our data to the dataset
	dims := make([]netcdf.Dim, 2)
	ht, wd := 5, 4
	dims[0], err = ds.AddDim("height", uint64(ht))
	if err != nil {
		return err
	}
	dims[1], err = ds.AddDim("width", uint64(wd))
	if err != nil {
		return err
	}

	// Add the variable to the dataset that will store our data
	v, err := ds.AddVar("gopher", netcdf.UBYTE, dims)
	if err != nil {
		return err
	}

	// Add a _FillValue to the variable's attributes
	// From C++ netCDF documentation:
	//   With netCDF-4 files, nc_put_att will notice if you are writing a _FillValue attribute,
	//   and will tell the HDF5 layer to use the specified fill value for that variable. With
	//   either classic or netCDF-4 files, a _FillValue attribute will be checked for validity,
	//   to make sure it has only one value and that its type matches the type of the associated
	//   variable.
	if err := v.Attr("_FillValue").WriteUint8s([]uint8{255}); err != nil {
		return err
	}

	// Add an attribute to the variable
	if err := v.Attr("year").WriteInt32s([]int32{2012}); err != nil {
		return err
	}

	// Create the data with the above dimensions and write it to the file.
	gopher := make([]uint8, ht*wd)
	i := 0
	for y := 0; y < ht; y++ {
		for x := 0; x < wd; x++ {
			gopher[i] = uint8(x + y)
			i++
		}
	}
	return v.WriteUint8s(gopher)
}

// ReadExampleFile reads the data in NetCDF file at filename and prints it out.
func ReadExampleFile(filename string) error {
	// Open example file in read-only mode. The dataset is returned.
	ds, err := netcdf.OpenFile(filename, netcdf.NOWRITE)
	if err != nil {
		return err
	}
	defer ds.Close()

	// Get the variable containing our data and read the data from the variable.
	v, err := ds.Var("gopher")
	if err != nil {
		return err
	}

	// Print variable attribute
	year, err := netcdf.GetInt32s(v.Attr("year"))
	if err != nil {
		return err
	}
	fmt.Printf("year = %v\n", year[0])

	// Read data from variable
	gopher, err := netcdf.GetUint8s(v)
	if err != nil {
		return err
	}

	// Get the length of the dimensions of the data.
	dims, err := v.LenDims()
	if err != nil {
		return err
	}

	// Print out the data
	i := 0
	for y := 0; y < int(dims[0]); y++ {
		for x := 0; x < int(dims[1]); x++ {
			fmt.Printf(" %d", gopher[i])
			i++
		}
		fmt.Printf("\n")
	}
	return nil
}

func Example() {
	// Create example file
	filename := "gopher.nc"
	if err := CreateExampleFile(filename); err != nil {
		log.Fatalf("creating example file failed: %v\n", err)
	}

	// Open and read example file
	if err := ReadExampleFile(filename); err != nil {
		log.Fatalf("reading example file failed: %v\n", err)
	}

	// Output:
	//  year = 2012
	//  0 1 2 3
	//  1 2 3 4
	//  2 3 4 5
	//  3 4 5 6
	//  4 5 6 7
}
