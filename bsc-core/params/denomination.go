// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// These are the multipliers for GOLDPN denominations.
// But actually these are the multipliers for Ethereum denominations. (for less maintenance and compatibility with other Ethereum-based projects, we kept the same names)
// Example: to convert 1 GWei to Wei:
//
//     new(big.Int).Mul(value, big.NewInt(params.GWei))

const (
    Wei   = 1
    GWei  = 1e9
    Ether = 1e18
)



