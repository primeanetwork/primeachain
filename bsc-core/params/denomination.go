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

// These are the multipliers for PRIM denominations.
// Example: to convert 1 mPRIM to attoPRIM:
//
//     new(big.Int).Mul(value, big.NewInt(params.mPRIM))

const (
	// attoPRIM is the smallest denomination of PRIM.
	// 1 PRIM = 10^18 attoPRIM (similar to Wei for Ether)
	attoPRIM = 1 // Smallest unit, equivalent to Wei or Jager
	// kPRIM is 1000 attoPRIM. Not typically used directly in this const block.
	// mPRIM is 1,000,000,000 attoPRIM (equivalent to GWei for Ether)
	mPRIM = 1e9 // Equivalent to GWei (for gas prices)
	// PRIM is the main unit of currency for PrimeaChain.
	PRIM = 1e18 // Main unit, equivalent to Ether or BNB
)


