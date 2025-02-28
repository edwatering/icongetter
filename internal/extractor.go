/*
 * @Author: Edwater
 * @Description:
 * @File: extractor
 * @Version: 1.0.0
 * @Data: 2025/2/21 15:45
 */

package internal

type Extractor interface {
	ExtractIcon(path string) error
}
