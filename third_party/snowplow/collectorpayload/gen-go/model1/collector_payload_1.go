// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package model1

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//   - Schema
//   - IpAddress
//   - Timestamp
//   - Encoding
//   - Collector
//   - UserAgent
//   - RefererUri
//   - Path
//   - Querystring
//   - Body
//   - Headers
//   - ContentType
//   - Hostname
//   - NetworkUserId
type CollectorPayload struct {
	// unused fields # 1 to 99
	IpAddress string `thrift:"ipAddress,100" db:"ipAddress" json:"ipAddress"`
	// unused fields # 101 to 199
	Timestamp int64 `thrift:"timestamp,200" db:"timestamp" json:"timestamp"`
	// unused fields # 201 to 209
	Encoding string `thrift:"encoding,210" db:"encoding" json:"encoding"`
	// unused fields # 211 to 219
	Collector string `thrift:"collector,220" db:"collector" json:"collector"`
	// unused fields # 221 to 299
	UserAgent *string `thrift:"userAgent,300" db:"userAgent" json:"userAgent,omitempty"`
	// unused fields # 301 to 309
	RefererUri *string `thrift:"refererUri,310" db:"refererUri" json:"refererUri,omitempty"`
	// unused fields # 311 to 319
	Path *string `thrift:"path,320" db:"path" json:"path,omitempty"`
	// unused fields # 321 to 329
	Querystring *string `thrift:"querystring,330" db:"querystring" json:"querystring,omitempty"`
	// unused fields # 331 to 339
	Body *string `thrift:"body,340" db:"body" json:"body,omitempty"`
	// unused fields # 341 to 349
	Headers []string `thrift:"headers,350" db:"headers" json:"headers,omitempty"`
	// unused fields # 351 to 359
	ContentType *string `thrift:"contentType,360" db:"contentType" json:"contentType,omitempty"`
	// unused fields # 361 to 399
	Hostname *string `thrift:"hostname,400" db:"hostname" json:"hostname,omitempty"`
	// unused fields # 401 to 409
	NetworkUserId *string `thrift:"networkUserId,410" db:"networkUserId" json:"networkUserId,omitempty"`
	// unused fields # 411 to 31336
	Schema string `thrift:"schema,31337" db:"schema" json:"schema"`
}

func NewCollectorPayload() *CollectorPayload {
	return &CollectorPayload{}
}

func (p *CollectorPayload) GetSchema() string {
	return p.Schema
}

func (p *CollectorPayload) GetIpAddress() string {
	return p.IpAddress
}

func (p *CollectorPayload) GetTimestamp() int64 {
	return p.Timestamp
}

func (p *CollectorPayload) GetEncoding() string {
	return p.Encoding
}

func (p *CollectorPayload) GetCollector() string {
	return p.Collector
}

var CollectorPayload_UserAgent_DEFAULT string

func (p *CollectorPayload) GetUserAgent() string {
	if !p.IsSetUserAgent() {
		return CollectorPayload_UserAgent_DEFAULT
	}
	return *p.UserAgent
}

var CollectorPayload_RefererUri_DEFAULT string

func (p *CollectorPayload) GetRefererUri() string {
	if !p.IsSetRefererUri() {
		return CollectorPayload_RefererUri_DEFAULT
	}
	return *p.RefererUri
}

var CollectorPayload_Path_DEFAULT string

func (p *CollectorPayload) GetPath() string {
	if !p.IsSetPath() {
		return CollectorPayload_Path_DEFAULT
	}
	return *p.Path
}

var CollectorPayload_Querystring_DEFAULT string

func (p *CollectorPayload) GetQuerystring() string {
	if !p.IsSetQuerystring() {
		return CollectorPayload_Querystring_DEFAULT
	}
	return *p.Querystring
}

var CollectorPayload_Body_DEFAULT string

func (p *CollectorPayload) GetBody() string {
	if !p.IsSetBody() {
		return CollectorPayload_Body_DEFAULT
	}
	return *p.Body
}

var CollectorPayload_Headers_DEFAULT []string

func (p *CollectorPayload) GetHeaders() []string {
	return p.Headers
}

var CollectorPayload_ContentType_DEFAULT string

func (p *CollectorPayload) GetContentType() string {
	if !p.IsSetContentType() {
		return CollectorPayload_ContentType_DEFAULT
	}
	return *p.ContentType
}

var CollectorPayload_Hostname_DEFAULT string

func (p *CollectorPayload) GetHostname() string {
	if !p.IsSetHostname() {
		return CollectorPayload_Hostname_DEFAULT
	}
	return *p.Hostname
}

var CollectorPayload_NetworkUserId_DEFAULT string

func (p *CollectorPayload) GetNetworkUserId() string {
	if !p.IsSetNetworkUserId() {
		return CollectorPayload_NetworkUserId_DEFAULT
	}
	return *p.NetworkUserId
}
func (p *CollectorPayload) IsSetUserAgent() bool {
	return p.UserAgent != nil
}

func (p *CollectorPayload) IsSetRefererUri() bool {
	return p.RefererUri != nil
}

func (p *CollectorPayload) IsSetPath() bool {
	return p.Path != nil
}

func (p *CollectorPayload) IsSetQuerystring() bool {
	return p.Querystring != nil
}

func (p *CollectorPayload) IsSetBody() bool {
	return p.Body != nil
}

func (p *CollectorPayload) IsSetHeaders() bool {
	return p.Headers != nil
}

func (p *CollectorPayload) IsSetContentType() bool {
	return p.ContentType != nil
}

func (p *CollectorPayload) IsSetHostname() bool {
	return p.Hostname != nil
}

func (p *CollectorPayload) IsSetNetworkUserId() bool {
	return p.NetworkUserId != nil
}

func (p *CollectorPayload) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 31337:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField31337(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 100:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField100(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 200:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField200(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 210:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField210(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 220:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField220(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 300:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField300(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 310:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField310(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 320:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField320(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 330:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField330(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 340:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField340(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 350:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField350(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 360:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField360(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 400:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField400(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 410:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField410(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CollectorPayload) ReadField31337(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 31337: ", err)
	} else {
		p.Schema = v
	}
	return nil
}

func (p *CollectorPayload) ReadField100(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 100: ", err)
	} else {
		p.IpAddress = v
	}
	return nil
}

func (p *CollectorPayload) ReadField200(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 200: ", err)
	} else {
		p.Timestamp = v
	}
	return nil
}

func (p *CollectorPayload) ReadField210(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 210: ", err)
	} else {
		p.Encoding = v
	}
	return nil
}

func (p *CollectorPayload) ReadField220(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 220: ", err)
	} else {
		p.Collector = v
	}
	return nil
}

func (p *CollectorPayload) ReadField300(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 300: ", err)
	} else {
		p.UserAgent = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField310(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 310: ", err)
	} else {
		p.RefererUri = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField320(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 320: ", err)
	} else {
		p.Path = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField330(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 330: ", err)
	} else {
		p.Querystring = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField340(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 340: ", err)
	} else {
		p.Body = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField350(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Headers = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Headers = append(p.Headers, _elem0)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *CollectorPayload) ReadField360(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 360: ", err)
	} else {
		p.ContentType = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField400(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 400: ", err)
	} else {
		p.Hostname = &v
	}
	return nil
}

func (p *CollectorPayload) ReadField410(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 410: ", err)
	} else {
		p.NetworkUserId = &v
	}
	return nil
}

func (p *CollectorPayload) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "CollectorPayload"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField100(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField200(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField210(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField220(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField300(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField310(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField320(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField330(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField340(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField350(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField360(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField400(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField410(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField31337(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CollectorPayload) writeField100(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "ipAddress", thrift.STRING, 100); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 100:ipAddress: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.IpAddress)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ipAddress (100) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 100:ipAddress: ", p), err)
	}
	return err
}

func (p *CollectorPayload) writeField200(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "timestamp", thrift.I64, 200); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 200:timestamp: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.Timestamp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.timestamp (200) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 200:timestamp: ", p), err)
	}
	return err
}

func (p *CollectorPayload) writeField210(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encoding", thrift.STRING, 210); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 210:encoding: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Encoding)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.encoding (210) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 210:encoding: ", p), err)
	}
	return err
}

func (p *CollectorPayload) writeField220(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "collector", thrift.STRING, 220); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 220:collector: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Collector)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.collector (220) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 220:collector: ", p), err)
	}
	return err
}

func (p *CollectorPayload) writeField300(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetUserAgent() {
		if err := oprot.WriteFieldBegin(ctx, "userAgent", thrift.STRING, 300); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 300:userAgent: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.UserAgent)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.userAgent (300) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 300:userAgent: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField310(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetRefererUri() {
		if err := oprot.WriteFieldBegin(ctx, "refererUri", thrift.STRING, 310); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 310:refererUri: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.RefererUri)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.refererUri (310) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 310:refererUri: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField320(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetPath() {
		if err := oprot.WriteFieldBegin(ctx, "path", thrift.STRING, 320); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 320:path: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Path)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.path (320) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 320:path: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField330(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetQuerystring() {
		if err := oprot.WriteFieldBegin(ctx, "querystring", thrift.STRING, 330); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 330:querystring: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Querystring)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.querystring (330) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 330:querystring: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField340(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetBody() {
		if err := oprot.WriteFieldBegin(ctx, "body", thrift.STRING, 340); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 340:body: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Body)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.body (340) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 340:body: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField350(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetHeaders() {
		if err := oprot.WriteFieldBegin(ctx, "headers", thrift.LIST, 350); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 350:headers: ", p), err)
		}
		if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.Headers)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Headers {
			if err := oprot.WriteString(ctx, string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(ctx); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 350:headers: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField360(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetContentType() {
		if err := oprot.WriteFieldBegin(ctx, "contentType", thrift.STRING, 360); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 360:contentType: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.ContentType)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.contentType (360) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 360:contentType: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField400(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetHostname() {
		if err := oprot.WriteFieldBegin(ctx, "hostname", thrift.STRING, 400); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 400:hostname: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Hostname)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.hostname (400) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 400:hostname: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField410(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetNetworkUserId() {
		if err := oprot.WriteFieldBegin(ctx, "networkUserId", thrift.STRING, 410); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 410:networkUserId: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.NetworkUserId)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.networkUserId (410) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 410:networkUserId: ", p), err)
		}
	}
	return err
}

func (p *CollectorPayload) writeField31337(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "schema", thrift.STRING, 31337); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 31337:schema: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Schema)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.schema (31337) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 31337:schema: ", p), err)
	}
	return err
}

func (p *CollectorPayload) Equals(other *CollectorPayload) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.IpAddress != other.IpAddress {
		return false
	}
	if p.Timestamp != other.Timestamp {
		return false
	}
	if p.Encoding != other.Encoding {
		return false
	}
	if p.Collector != other.Collector {
		return false
	}
	if p.UserAgent != other.UserAgent {
		if p.UserAgent == nil || other.UserAgent == nil {
			return false
		}
		if (*p.UserAgent) != (*other.UserAgent) {
			return false
		}
	}
	if p.RefererUri != other.RefererUri {
		if p.RefererUri == nil || other.RefererUri == nil {
			return false
		}
		if (*p.RefererUri) != (*other.RefererUri) {
			return false
		}
	}
	if p.Path != other.Path {
		if p.Path == nil || other.Path == nil {
			return false
		}
		if (*p.Path) != (*other.Path) {
			return false
		}
	}
	if p.Querystring != other.Querystring {
		if p.Querystring == nil || other.Querystring == nil {
			return false
		}
		if (*p.Querystring) != (*other.Querystring) {
			return false
		}
	}
	if p.Body != other.Body {
		if p.Body == nil || other.Body == nil {
			return false
		}
		if (*p.Body) != (*other.Body) {
			return false
		}
	}
	if len(p.Headers) != len(other.Headers) {
		return false
	}
	for i, _tgt := range p.Headers {
		_src1 := other.Headers[i]
		if _tgt != _src1 {
			return false
		}
	}
	if p.ContentType != other.ContentType {
		if p.ContentType == nil || other.ContentType == nil {
			return false
		}
		if (*p.ContentType) != (*other.ContentType) {
			return false
		}
	}
	if p.Hostname != other.Hostname {
		if p.Hostname == nil || other.Hostname == nil {
			return false
		}
		if (*p.Hostname) != (*other.Hostname) {
			return false
		}
	}
	if p.NetworkUserId != other.NetworkUserId {
		if p.NetworkUserId == nil || other.NetworkUserId == nil {
			return false
		}
		if (*p.NetworkUserId) != (*other.NetworkUserId) {
			return false
		}
	}
	if p.Schema != other.Schema {
		return false
	}
	return true
}

func (p *CollectorPayload) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CollectorPayload(%+v)", *p)
}
