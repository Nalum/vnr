// VNR - Plugin for NewRelic which monitors `varnishstat`
// Copyright (C) 2015  Luke Mallon
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package structs

import "fmt"

type Config struct {
	Key         string
	Instances   []string
	Interval    int
	VarnishStat VarnishStat
}

type VarnishStat struct {
	Type string
	Path string
}

func (c Config) String() string {
	return fmt.Sprintf("Key: %v - Instances: %v - Interval: %v - VarnishStat.Type: %v - VarnishStat.Path: %v", c.Key, c.Instances, c.Interval, c.VarnishStat.Type, c.VarnishStat.Path)
}
